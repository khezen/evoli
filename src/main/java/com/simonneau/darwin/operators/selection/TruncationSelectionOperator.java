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
import java.util.Iterator;

/**
 *
 * @author simonneau
 */
public class TruncationSelectionOperator extends SelectionOperator {

    private static TruncationSelectionOperator instance;
    private static String LABEL = "Truncation selection";

    private TruncationSelectionOperator() {
        super(LABEL);
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
    public Population buildNextGeneration(Population population, int survivorSize) {

        Population nextPopulation = new Population(population.getObservableVolume());

        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population.getIndividuals());
            
        } else {
            population.sort();
            Iterator<Individual> iterator = population.iterator();
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
