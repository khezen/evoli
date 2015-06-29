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
import java.util.LinkedList;

/**
 *
 * @author simonneau
 */
public class TournamentSelectionOperator implements SelectionOperator {

    private static TournamentSelectionOperator instance;
    private static final String LABEL = "Tournament selection";
    private LinkedList<Genotype> draft;

    private TournamentSelectionOperator() {
    }

    /**
     *
     * @return
     */
    public static TournamentSelectionOperator getInstance() {
        if (TournamentSelectionOperator.instance == null) {
            instance = new TournamentSelectionOperator();
        }
        return instance;
    }

    /**
     * select survivorSize individuals form population. each individuals selected win a tournament.
     * @param population
     * @param survivorSize
     * @return
     */
    @Override
    public Population<? extends Genotype> buildNextGeneration(Population<? extends Genotype> population, int survivorSize) {
        PopulationImpl nextPopulation = new PopulationImpl(population.getPopulationSize());
        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population);
        } else {
            ArrayList<Genotype> individuals = new ArrayList<>();
            this.draft = new LinkedList<>();
            this.draft.addAll(population);
            int survivorCount = 0;
            int size;
            while (survivorCount < survivorSize) {
                individuals.clear();
                individuals.addAll(this.draft);
                this.draft.clear();
                size = individuals.size();
                if (size == 1) {
                    nextPopulation.add(individuals.get(0));
                    survivorCount++;
                }
                while (size > 1 && survivorCount < survivorSize) {
                    int firstChallengerIndex = (int) Math.round(Math.random() * (size - 1));
                    Genotype firstChallenger = individuals.remove(firstChallengerIndex);
                    size--;
                    double firstScore = firstChallenger.getSurvivalScore();
                    int secondChallengerIndex = (int) Math.round(Math.random() * (size - 1));
                    Genotype secondChallenger = individuals.remove(secondChallengerIndex);
                    size--;
                    double secondScore = secondChallenger.getSurvivalScore();
                    if (firstScore > secondScore) {
                        nextPopulation.add(firstChallenger);
                        survivorCount++;
                        this.draft.add(secondChallenger);
                    } else if (firstScore < secondScore) {
                        nextPopulation.add(secondChallenger);
                        survivorCount++;
                        this.draft.add(firstChallenger);
                    } else {
                        nextPopulation.add(firstChallenger);
                        survivorCount++;
                        if (survivorCount < survivorSize) {
                            nextPopulation.add(secondChallenger);
                            survivorCount++;
                        }
                    }
                }
                if (size == 1) {
                    this.draft.add(individuals.get(0));
                }
            }
        }
        return nextPopulation;
    }
}
