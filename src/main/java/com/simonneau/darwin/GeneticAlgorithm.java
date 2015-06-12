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


import com.simonneau.darwin.population.Population;
import com.simonneau.darwin.problem.Problem;
import java.util.Collection;
import java.util.LinkedList;

/**
 * Hello world!
 *
 */
public class GeneticAlgorithm  {

    private LinkedList<Problem> problems;
    private Problem selectedProblem;
    private GeneticEngine geneticEngine;
    private boolean UIvisible = false;

    /**
     *
     */
    public GeneticAlgorithm() {
        this.problems = new LinkedList<>();
        this.init();

    }

    /**
     *
     * @return
     */
    public Problem getSelectedProblem() {
        return selectedProblem;
    }

    /**
     * set the selected problem with selectedProblem.
     *
     * @param SelectedProblem
     */
    public void setSelectedProblem(Problem SelectedProblem) {
        this.selectedProblem = SelectedProblem;
        this.restart();
    }

    /**
     * add a problem to 'this'.
     *
     * @param problem
     */
    public final void addProblem(Problem problem) {
        this.problems.add(problem);
        this.selectedProblem = problem;
    }

    /**
     * add all problems from foreignproblems to 'this'.
     *
     * @param foreignProblems
     */
    public void addAll(Collection<Problem> foreignProblems) {
        foreignProblems.stream().forEach((pb) -> {
            this.addProblem(pb);
        });
    }

    /**
     * return 'this' problems.
     *
     * @return 'this' available problems.
     */
    public LinkedList<Problem> getProblems() {
        return this.problems;
    }

    /**
     *
     * @param index
     * @return the problem from index 'index in 'this' availables problems.
     */
    public Problem getProblem(int index) {
        return this.problems.get(index);
    }

    /**
     * run the genetic algorithm.
     */
    private void init() {
        if (this.geneticEngine == null) {
            this.geneticEngine = new GeneticEngine(this.selectedProblem);
        } else {
            this.geneticEngine.setProblem(selectedProblem);
        }
    }


    /**
     * pause the engine.
     */
    public void pause() {
        if (this.geneticEngine != null && !this.geneticEngine.isPaused()) {
            this.geneticEngine.pause();
        }
    }

    /**
     * run the engine.
     */
    public void run() {
        if (this.geneticEngine != null && this.geneticEngine.isPaused()) {
            this.geneticEngine.resume();
        }
    }

    /**
     * return the currentPopulation.
     *
     * @return
     */
    public Population getCurrentPopulation() {
        if (this.geneticEngine != null) {
            return this.geneticEngine.getPopulation();
        } else {
            throw new NeitherSelectedProblemException();

        }
    }

    /**
     * restart the engine.
     */
    public void restart() {
        this.geneticEngine.setProblem(this.selectedProblem);
    }

    /**
     * quit the geneticAlgorithm application.
     */
    public void quit() {
        this.geneticEngine.pause();
        
    }

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        GeneticAlgorithm ga = new GeneticAlgorithm();
    }
}
