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
import com.simonneau.darwin.population.IndividualImpl;
import com.simonneau.darwin.population.Population;
import com.simonneau.darwin.population.PopulationImpl;
import java.util.Iterator;

/**
 *
 * @author simonneau
 */
public class TruncationSelectionOperator implements SelectionOperator {

    private static TruncationSelectionOperator instance;
    private static final String LABEL = "Truncation selection";

    private TruncationSelectionOperator() {
    }

    /**
     *
     * @return
     */
    public static TruncationSelectionOperator getInstance() {
        if (TruncationSelectionOperator.instance == null) {
            instance = new TruncationSelectionOperator();
        }
        return instance;
    }

    /**
     * select survivorSize individuals form population. select the survivorSize best individuals.
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
            population.sort();
            Iterator<? extends Individual> iterator = population.iterator();
            Individual individual;
            int i = 0;
            while (iterator.hasNext() && i < survivorSize) {
                individual = iterator.next();
                nextPopulation.add(individual);
                i++;
            }
        }
        return nextPopulation;
    }
}
