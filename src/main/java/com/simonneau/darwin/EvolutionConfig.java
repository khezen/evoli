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


import com.simonneau.darwin.operators.CrossOverOperator;
import com.simonneau.darwin.operators.EvaluationOperator;
import com.simonneau.darwin.operators.MutationOperator;
import com.simonneau.darwin.operators.OperatorSet;
import com.simonneau.darwin.operators.ProportionalPerfomanceSelectionOperator;
import com.simonneau.darwin.operators.ProportionalRankingSelectionOperator;
import com.simonneau.darwin.operators.RandomSelectionOperator;
import com.simonneau.darwin.operators.SelectionOperator;
import com.simonneau.darwin.operators.TournamentSelectionOperator;
import com.simonneau.darwin.operators.TruncationSelectionOperator;
import com.simonneau.darwin.population.IndividualImpl;

import java.util.LinkedList;

/**
 *
 * @param <T> 
 * @author simonneau
 */
public class EvolutionConfig<T extends IndividualImpl> {

    private LinkedList<MutationOperator<T>> availableMutationOperators = new LinkedList<>();
    private LinkedList<CrossOverOperator<T>> availableCrossOverOperators = new LinkedList<>();
    private LinkedList<SelectionOperator> availableSelectionOperators = new LinkedList<>();
    private LinkedList<EvaluationOperator<T>> availableEvaluationOperator = new LinkedList<>();
    
    private int populationSize = 20;
    private String label;
    private double mutationProbability = 0.1;
    private double crossProbability = 0.2;
    private OperatorSet operators = new OperatorSet();
    private StopCriteria stopCriteria;

    /**
     *
     */
    public EvolutionConfig() {
        this.stopCriteria = new StopCriteria();
        this.addSelectionOperator(RandomSelectionOperator.getInstance());
        this.addSelectionOperator(TruncationSelectionOperator.getInstance());
        this.addSelectionOperator(TournamentSelectionOperator.getInstance());
        this.addSelectionOperator(ProportionalPerfomanceSelectionOperator.getInstance());
        this.addSelectionOperator(ProportionalRankingSelectionOperator.getInstance());
    }

    /**
     *
     * @return
     */
    public StopCriteria getStopCriteria() {
        return this.stopCriteria;
    }

    /**
     *
     * @return the timeout stop criterion.
     */
    public long getTimeout() {
        return stopCriteria.getTimeout();
    }

    /**
     *
     * @return
     */
    public LinkedList<MutationOperator<T>> getAvailableMutationOperators() {
        return availableMutationOperators;
    }

    /**
     *
     * @return
     */
    public LinkedList<CrossOverOperator<T>> getAvailableCrossOverOperators() {
        return availableCrossOverOperators;
    }

    /**
     *
     * @return
     */
    public LinkedList<SelectionOperator> getAvailableSelectionOperators() {
        return availableSelectionOperators;
    }

    /**
     *
     * @return
     */
    public LinkedList<EvaluationOperator<T>> getAvailableEvaluationOperator() {
        return availableEvaluationOperator;
    }

    /**
     *
     * @return the max step count stop criterion.
     */
    public long getMaxStepCount() {
        return stopCriteria.getMaxStepCount();
    }

    /**
     *
     * @return 'this' label.
     */
    public String getLabel() {
        return label;
    }

    /**
     *
     * @return 'this' mutation probability.
     */
    public double getMutationProbability() {
        return mutationProbability;
    }

    /**
     *
     * @return 'this' cross probability.
     */
    public double getCrossProbability() {
        return crossProbability;
    }

    /**
     *
     * @return 'this' selected MutationOperator.
     */
    public MutationOperator<T> getSelectedMutationOperator() {
        return operators.getMutationOperator();
    }

    /**
     *
     * @return 'this' selected CrossOverOperator.
     */
    public CrossOverOperator<T> getSelectedCrossOverOperator() {
        return operators.getCrossoverOperator();
    }

    /**
     *
     * @return 'this.selected SelectionOperator.
     */
    public SelectionOperator getSelectedSelectionOperator() {
        return operators.getSelectionOperator();
    }

    /**
     *
     * @return 'this' selected EvaluationOperator.
     */
    public EvaluationOperator<T> getSelectedEvaluationOperator() {
        return operators.getEvaluationOperator();
    }

    /**
     *
     * @param mutationProbability
     */
    public void setMutationProbability(double mutationProbability) {
        this.mutationProbability = mutationProbability;
    }

    /**
     *
     * @param crossProbability
     */
    public void setCrossProbability(double crossProbability) {
        this.crossProbability = crossProbability;
    }

    /**
     *
     * @param selectedMutationOperator
     */
    public void setSelectedMutationOperator(MutationOperator<T> selectedMutationOperator) {
        this.operators.setMutationOperator(selectedMutationOperator);
    }

    /**
     *
     * @param selectedCrossOverOperation
     */
    public void setSelectedCrossOverOperation(CrossOverOperator<T> selectedCrossOverOperation) {
        this.operators.setCrossoverOperator(selectedCrossOverOperation);
    }

    /**
     *
     * @param selectedSelectionOperator
     */
    public void setSelectedSelectionOperator(SelectionOperator selectedSelectionOperator) {
        this.operators.setSelectionOperator(selectedSelectionOperator);
    }

    /**
     *
     * @param selectedEvaluationOperator
     */
    public void setSelectedEvaluationOperator(EvaluationOperator<T> selectedEvaluationOperator) {
        this.operators.setEvaluationOperator(selectedEvaluationOperator);
    }

    /**
     *
     * @param populationSize
     */
    public void setPopulationSize(int populationSize) {
        this.populationSize = populationSize;
    }

    /**
     *
     * @param label
     */
    public void setLabel(String label) {
        this.label = label;
    }

    /**
     *
     * @return 'this' populatio size.
     */
    public int getPopulationSize() {
        return populationSize;
    }
    
    /**
     *
     * @param operators
     */
    protected void setSelectedOperators(OperatorSet operators) {
        this.operators = operators;
    }

    

    /**
     *
     * @param operator
     */
    public void addMutationOperator(MutationOperator<T> operator) {
        this.availableMutationOperators.add(operator);
        this.setSelectedMutationOperator(operator);
    }

    /**
     *
     * @param operator
     */
    public void addCrossOverOperator(CrossOverOperator<T> operator) {
        this.availableCrossOverOperators.add(operator);
        this.setSelectedCrossOverOperation(operator);
    }

    /**
     *
     * @param operator
     */
    public final void addSelectionOperator(SelectionOperator operator) {
        this.availableSelectionOperators.add(operator);
        this.setSelectedSelectionOperator(operator);
    }

    /**
     *
     * @param operator
     */
    public void addEvaluationOperator(EvaluationOperator<T> operator) {
        this.availableEvaluationOperator.add(operator);
        this.setSelectedEvaluationOperator(operator);
    }

    /**
     *
     * @return
     */
    public double getEvolutionCriterion() {
        return this.stopCriteria.getEvolutionCriterion();
    }

    /**
     *
     * @param stepCount
     * @param time
     * @param evolutionCoeff
     * @return return true if the stop criteria are reached. false other wise.
     */
    public boolean stopCriteriaAreReached(long stepCount, long time, double evolutionCoeff) {
        return this.stopCriteria.areReached(stepCount, time, evolutionCoeff);
    }

    /**
     *
     * @return 'this' label.
     */
    @Override
    public final String toString() {
        return this.getLabel();
    }
}
