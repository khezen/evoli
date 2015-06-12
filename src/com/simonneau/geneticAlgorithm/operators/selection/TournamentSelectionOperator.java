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

import com.simonneau.geneticAlgorithm.population.Individual;
import com.simonneau.geneticAlgorithm.population.Population;
import java.util.ArrayList;
import java.util.LinkedList;

/**
 *
 * @author simonneau
 */
public class TournamentSelectionOperator extends SelectionOperator {

    private static TournamentSelectionOperator instance;
    private static String LABEL = "Tournament selection";
    private LinkedList<Individual> draft;

    private TournamentSelectionOperator() {
        super(LABEL);
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
    public Population buildNextGeneration(Population population, int survivorSize) {

        Population nextPopulation = new Population(population.getObservableVolume());

        if (population.size() <= survivorSize) {
            nextPopulation.addAll(population.getIndividuals());

        } else {

            ArrayList<Individual> individuals = new ArrayList<>();
            this.draft = new LinkedList<>();
            this.draft.addAll(population.getIndividuals());



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
                    Individual firstChallenger = individuals.remove(firstChallengerIndex);
                    size--;
                    double firstScore = firstChallenger.getScore();

                    int secondChallengerIndex = (int) Math.round(Math.random() * (size - 1));
                    Individual secondChallenger = individuals.remove(secondChallengerIndex);
                    size--;
                    double secondScore = secondChallenger.getScore();

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
