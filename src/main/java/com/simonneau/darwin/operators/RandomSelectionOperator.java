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
import java.util.LinkedList;

/**
 *
 * @author simonneau
 */
public class RandomSelectionOperator implements SelectionOperator {

    private static RandomSelectionOperator instance;
    private static final String LABEL = "Random selection";

    private RandomSelectionOperator() {
    }

    public static RandomSelectionOperator getInstance() {
        if (RandomSelectionOperator.instance == null) {
            instance = new RandomSelectionOperator();
        }
        return instance;
    }

    /**
     * select survivorSize individuals randomly from population.
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
            LinkedList<Genotype> individuals = new LinkedList<>(population);
            int survivorCount = 0;
            int size = individuals.size();
            while (survivorCount < survivorSize) {
                int index = (int) Math.round(Math.random() * (size - 1));
                nextPopulation.add(individuals.remove(index));
                survivorCount++;
                size--;
            }
        }
        return nextPopulation;
    }
}
