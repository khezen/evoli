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
package com.simonneau.geneticAlgorithm.operators.selection;

import com.simonneau.geneticAlgorithm.operators.Operator;
import com.simonneau.geneticAlgorithm.population.Population;

/**
 *
 * @author simonneau
 */
public abstract class SelectionOperator extends Operator {
    
    /**
     *
     * @param label
     */
    public SelectionOperator(String label){
        super(label);
    }

    /**
     * select survivorSize individuals form population.
     * @param population
     * @param survivorSize
     * @return
     */
    public abstract Population buildNextGeneration(Population population, int survivorSize);
}
