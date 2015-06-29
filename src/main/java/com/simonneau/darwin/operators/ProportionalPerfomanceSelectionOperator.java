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
package com.simonneau.darwin.operators;

import com.simonneau.darwin.population.Genotype;
import com.simonneau.darwin.population.Population;
import com.simonneau.darwin.population.PopulationImpl;
import java.util.ArrayList;
import java.util.Collections;

/**
 *
 * @author simonneau
 */
public class ProportionalPerfomanceSelectionOperator implements SelectionOperator {

    private static ProportionalPerfomanceSelectionOperator instance;
    public static final String LABEL = "Proportional perfomance selection";
    private ArrayList<Double> scores;

    private ProportionalPerfomanceSelectionOperator() {
    }

    /**
     *
     * @return
     */
    public static ProportionalPerfomanceSelectionOperator getInstance() {
        if (ProportionalPerfomanceSelectionOperator.instance == null) {
            instance = new ProportionalPerfomanceSelectionOperator();
        }
        return instance;
    }

    /**
     * select survivorSize individuals from population. each individuals have a
     * cvhance to survivre proportional with his performance.
     *
     * @param population
     * @param survivorSize
     * @return
     */
    @Override
    public Population<? extends Genotype> buildNextGeneration(Population<? extends Genotype> population, int survivorSize) {
        Population nextPopulation = new PopulationImpl(population.getPopulationSize());
        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population);
        } else {
            double totalScore = this.getTotalScore(population);
            Population p = population.clone();
            double score;
            int survivorCount = 0;
            int i;
            int size;
            while (survivorCount < survivorSize) {
                i = 0;
                size = p.size();
                while (i < size && survivorCount < survivorSize) {
                    score = this.scores.get(i);
                    if (Math.random() <= score / totalScore) {
                        nextPopulation.add(p.remove(i));
                        this.scores.remove(i);
                        size--;
                        totalScore -= score;
                        survivorCount++;
                    }
                    i++;
                }
            }
        }
        return nextPopulation;
    }

    private double getTotalScore(Population<? extends Genotype> population) {
        this.scores = new ArrayList<>(population.size());
        double minScore = this.getminScore(population);
        double totalScore = 0;
        double score;
        for (Genotype individual : population) {
            score = individual.getSurvivalScore() - minScore + 1;
            this.scores.add(score);
            totalScore += score;
        }
        return totalScore;
    }

    private double getminScore(Population<? extends Genotype> population) {
        return Collections.min(population).getSurvivalScore();
    }
}
