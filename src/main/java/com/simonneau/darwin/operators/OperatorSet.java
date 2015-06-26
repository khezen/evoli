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

/**
 *
 * @author simonneau
 */
public class OperatorSet {
    
    private CrossOverOperator crossoverOperator;
    private EvaluationOperator evaluationOperator;
    private MutationOperator mutationOperator;
    private SelectionOperator selectionOperator;
    

    /**
     *
     * @return 'this' selected CrossoverOperator;
     */
    public CrossOverOperator getCrossoverOperator() {
        return crossoverOperator;
    }

    /**
     *
     * @return 'this' selected EvaluationOperator.
     */
    public EvaluationOperator getEvaluationOperator() {
        return evaluationOperator;
    }

    /**
     *
     * @return 'this' selected MutationOperator.
     */
    public MutationOperator getMutationOperator() {
        return mutationOperator;
    }

    /**
     *
     * @return 'this' selected SelectionOperator.
     */
    public SelectionOperator getSelectionOperator() {
        return selectionOperator;
    }

    /**
     * set 'this' selected CrossoverOperator.
     * @param crossoverOperator
     */
    public void setCrossoverOperator(CrossOverOperator crossoverOperator) {
        this.crossoverOperator = crossoverOperator;
    }

    /**
     * set this' selected EvaluationOperator.
     * @param evaluationOperator
     */
    public void setEvaluationOperator(EvaluationOperator evaluationOperator) {
        this.evaluationOperator = evaluationOperator;
    }

    /**
     * set 'this' selected MutationOperator.
     * @param mutationOperator
     */
    public void setMutationOperator(MutationOperator mutationOperator) {
        this.mutationOperator = mutationOperator;
    }

    /**
     * set 'this' selected SelectionOperator.
     * @param selectionOperator
     */
    public void setSelectionOperator(SelectionOperator selectionOperator) {
        this.selectionOperator = selectionOperator;
    }
    
    
    
    
}
