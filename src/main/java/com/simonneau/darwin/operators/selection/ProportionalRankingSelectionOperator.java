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
package com.simonneau.darwin.operators.selection;

import com.simonneau.darwin.population.Individual;
import com.simonneau.darwin.population.Population;
import java.util.ArrayList;
import java.util.List;

/**
 *
 * @author simonneau
 */
public class ProportionalRankingSelectionOperator extends SelectionOperator {

    private static ProportionalRankingSelectionOperator instance;
    private static String LABEL = "Proportional ranking selection";
    private ArrayList<Integer> ranking;
    private int poolRange;

    private ProportionalRankingSelectionOperator() {
        super(LABEL);
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
    public Population buildNextGeneration(Population population, int survivorSize) {

        Population nextPopulation = new Population(population.getObservableVolume());

        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population.getIndividuals());

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

    private void madeRanking(Population population) {

        this.ranking = new ArrayList<>(population.size());
        this.poolRange = 0;
        population.sort();

        List<Individual> individuals = population.getIndividuals();
        int size = population.size();

        for (int i = 1; i <= size; i++) {

            this.ranking.add(i);
            this.poolRange += i;
        }
    }
}
