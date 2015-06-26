/*
 * Copyright (C) 2015 Guillaume Simonneau, simonneaug@gmail.com
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package com.simonneau.darwin;


import com.simonneau.darwin.util.Chronometer;
import com.simonneau.darwin.operators.CrossOverOperator;
import com.simonneau.darwin.operators.MutationOperator;
import com.simonneau.darwin.population.Individual;
import com.simonneau.darwin.population.Population;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.LinkedList;

/**
 *
 * @author simonneau
 * @param <T>
 */
public class EvolutionEngine<T extends Individual> implements Runnable {

    private boolean pause = true;
    private long stepCount = 0;
    private double evolutionCriterion = 1;
    private double previousBestScore = 0;
    private boolean firstStepDone = false;
    private Chronometer chronometer;
    private Thread engine;
    
    private Population<T> population;
    private EvolutionConfig config;
    private final PopulationFactory<T> populationFactory;
    

    /**
     *
     * @param problem
     * @param populationFactory
     */
    public EvolutionEngine(EvolutionConfig problem, PopulationFactory populationFactory) {
        this.chronometer = new Chronometer();
        this.setProblem(problem);
        this.populationFactory = populationFactory;
    }

    /**
     *
     * @return true if the engine is paused. false otherwise.
     */
    public boolean isPaused() {
        return this.pause;
    }

    public void resizePop() {
        Population<T> pop = this.populationFactory.createRandom();
        Iterator<T> it = pop.iterator();
        int newSize = this.config.getPopulationSize();
        int currentSize = this.population.size();
        while (it.hasNext() && currentSize < newSize) {
            this.population.add(it.next());
            currentSize++;
        }
    }

    public void refreshPopulation() {
        if (!this.pause) {
            this.pause();
            this.resume();
        }
    }

    public void reset() {
        this.pause();
        this.chronometer.reset();
        this.evolutionCriterion = 1;
        this.stepCount = 0;
        this.setPopulation(this.populationFactory.createRandom());
    }

    public Population getPopulation() {
        return population;
    }

    /**
     *
     * @return the treated config.
     */
    public EvolutionConfig getProblem() {
        return config;
    }

    /**
     *
     * @param problem
     */
    public final void setProblem(EvolutionConfig problem) {
        this.config = problem;
        this.pause = true;
        this.reset();
    }

    /**
     *
     * @return
     */
    public long getStepCount() {
        return this.stepCount;
    }

    private void setPopulation(Population<T> population) {
        this.population = population;
        this.firstStepDone = false;
        this.evaluationStep();
    }

    /**
     * resume this.
     */
    public void resume() {
        if (this.pause) {
            this.pause = false;
            this.chronometer.start();
            this.engine();
        }
    }

    /**
     * pause this.
     */
    public void pause() {
        if (!this.pause) {
            this.pause = true;
            this.chronometer.stop();
            if (this.engine != null) {
                try {
                    this.engine.join();
                } catch (InterruptedException e) {
                    //TODO
                }
            }
        }
    }

    /**
     * run the engine in an independant thread.
     */
    @Override
    public void run() {
        while (!this.pause && !this.config.stopCriteriaAreReached(this.stepCount, this.chronometer.getTime(), this.evolutionCriterion)) {
            this.evolve();
        }
        this.pause();
    }
    
    private void engine() {
        this.engine = new Thread(this);
        engine.start();
    }
    
    public T getBestSolution() {
        return this.population.getAlphaIndividual();
    }

    /**
     * Evaluate all the individuals using their evaluation method.
     */
    private void evaluationStep() {

        this.population.stream().forEach((individual) -> {
            this.config.getSelectedEvaluationOperator().evaluate(individual);
        });

        this.population.sort();
        this.computeEvolutionCriterion();
    }

    private void computeEvolutionCriterion() {
        double bestScore = this.population.getAlphaIndividual().getSurvivalScore();
        if (this.firstStepDone) {
            this.evolutionCriterion = Math.abs(this.previousBestScore - bestScore) / Math.abs(this.previousBestScore);
        } else {
            this.firstStepDone = true;
        }
        this.previousBestScore = bestScore;
    }

    /**
     * Selects the survivals of the current generation using the selected
     * selection operator.
     */
    private void buildNextGeneration() {

        this.evaluationStep();
        Population<T> pop = (Population<T>) this.config.getSelectedSelectionOperator().buildNextGeneration(this.population, this.config.getPopulationSize());
        this.population.clear();
        this.population.addAll(pop);
    }

    /**
     * Randomly crosses Individuals between them using the selected cross over
     * operator.
     */
    private void crossOverStep() {
        LinkedList<T> crossQueue = new LinkedList<>();
        CrossOverOperator<T> crossoverOperator = this.config.getSelectedCrossOverOperator();
        this.population.stream().filter((individual) -> (Math.random() < this.config.getCrossProbability())).forEach((individual) -> {
            crossQueue.add(individual);
        });
        int queueSize = crossQueue.size();
        T male;
        T female = null;
        int nbCandidates;
        double sexAppeal;
        while (queueSize > 1) {
            male = crossQueue.remove(0);
            queueSize--;
            nbCandidates = queueSize;
            boolean done = false;
            while (nbCandidates > 0 && !done) {
                Iterator<T> solutionIterator = crossQueue.iterator();
                female = solutionIterator.next();
                sexAppeal = 1 / nbCandidates;
                if (Math.random() < sexAppeal) {
                    solutionIterator.remove();
                    queueSize--;
                    done = true;
                } else {
                    nbCandidates--;
                }
            }
            this.population.add(crossoverOperator.cross(male, female));
        }
    }

    /**
     * Randomly makes some individual victim of mutations using the selected
     * mutation operator.
     */
    private void mutationStep() {
        MutationOperator<T> mutationOperator = this.config.getSelectedMutationOperator();
        ArrayList<T> mutants = new ArrayList<>();
        this.population.stream().filter((individual) -> (Math.random() < this.config.getMutationProbability())).forEach((individual) -> {
            mutants.add(mutationOperator.mutate(individual));
        });
        this.population.addAll(mutants);
    }

    /**
     * process all the steps of genetic algorithms.
     */
    public void evolve() {
        this.crossOverStep();
        this.mutationStep();
        this.buildNextGeneration();
        this.stepCount++;
    }

    /**
     * process only one step generation.
     */
    public void step() {
        if (!this.config.stopCriteriaAreReached(this.stepCount, this.chronometer.getTime(), this.evolutionCriterion)) {
            this.chronometer.start();
            this.evolve();
            this.chronometer.stop();          
        }
    }
}
