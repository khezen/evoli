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


import com.simonneau.darwin.population.Individual;
import com.simonneau.darwin.population.Population;
import com.simonneau.darwin.population.PopulationImpl;
import java.util.ArrayList;


/**
 *
 * @author simonneau
 */
public class ProportionalRankingSelectionOperator implements SelectionOperator {

    private static ProportionalRankingSelectionOperator instance;
    private static final String LABEL = "Proportional ranking selection";
    private ArrayList<Integer> ranking;
    private int poolRange;

    private ProportionalRankingSelectionOperator() {
    }

    /**
     *
     * @return
     */
    public static ProportionalRankingSelectionOperator getInstance() {
        if (ProportionalRankingSelectionOperator.instance == null) {
            instance = new ProportionalRankingSelectionOperator();
        }
        return instance;
    }

    /**
     * select survivorSize individuals from population. each individuals have a cvhance to survivre proportional with his rank.
     * @param population
     * @param survivorSize
     * @return
     */
    @Override
    public Population<? extends Individual> buildNextGeneration(Population<? extends Individual> population, int survivorSize) {

        PopulationImpl nextPopulation = new PopulationImpl(population.getPopulationSize());
        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population);
        } else {
            this.madeRanking(population);
            Population p = population.clone();
            int survivorCount = 0;
            int i;
            int size;
            int initialSize = p.size();
            double adaptability;
            while (survivorCount < survivorSize ) {
                i = 0;
                size = p.size();
                while (i < size && survivorCount < survivorSize) {
                    adaptability = initialSize - this.ranking.get(i) + 1;
                    if (Math.random() <= adaptability / this.poolRange) {

                        nextPopulation.add(p.remove(i));
                        this.ranking.remove(i);
                        this.poolRange -= adaptability;
                        size--;
                        survivorCount++;
                    }
                    i++;
                }
            }
        }
        return nextPopulation;
    }

    private void madeRanking(Population<? extends Individual> population) {
        this.ranking = new ArrayList<>(population.size());
        this.poolRange = 0;
        population.sort();
        int size = population.size();
        for (int i = 1; i <= size; i++) {
            this.ranking.add(i);
            this.poolRange += i;
        }
    }
}
