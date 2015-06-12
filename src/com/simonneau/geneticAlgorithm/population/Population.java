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
package com.simonneau.geneticAlgorithm.population;

import com.simonneau.geneticAlgorithm.population.Individual;
import com.simonneau.geneticAlgorithm.population.IndividualComparator;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.Iterator;
import java.util.LinkedList;

/**
 *
 * @author simonneau
 */
public class Population {

    private ArrayList<Individual> individuals;
    private int observableVolume = 1;
    private boolean semaphoreAccess = true;
    
    private Individual alphaIndividual;

    /**
     *
     */
    public Population() {
        this(1);
    }

    /**
     *
     * @param observableVolume
     */
    public Population(int observableVolume) {
        this.individuals = new ArrayList<>();
        this.observableVolume = observableVolume;
    }

    /**
     * Adds an individual to the 'this' Population.
     *
     * @param s
     */
    public void add(Individual s) {
        this.individuals.add(s);
    }
    
    /**
     * adds all individuals from coll to 'this'.
     * @param coll
     */
    public void addAll(Collection<? extends Individual> coll) {
        this.individuals.addAll(coll);
    }

    /**
     * adds all individuals from coll begining to the index 'index' to 'this'.
     * @param index
     * @param coll
     */
    public void addAll(int index, Collection<? extends Individual> coll) {
        this.individuals.addAll(index, coll);
    }

    /**
     *
     * @return an iterator on individuals.
     */
    public Iterator<Individual> iterator() {
        return this.individuals.iterator();
    }

    /**
     *
     * @return 'this' individuals.
     */
    public ArrayList<Individual> getIndividuals() {
        return individuals;
    }

    /**
     *
     * @param index
     * @return
     */
    public Individual get(int index) {
        return this.individuals.get(index);
    }

    /**
     * set the obserable volume.
     * @param observableVolume
     */
    public void setObservableVolume(int observableVolume) {
        int size = this.individuals.size();
        if (observableVolume > size) {
            observableVolume = size;
        } else if (observableVolume < 1) {
            observableVolume = 1;
        }
        this.observableVolume = observableVolume;
    }

    /**
     *
     * @return
     */
    public int getObservableVolume() {
        return observableVolume;
    }

    /**
     *
     * @param population
     */
    public void setIndividuals(Population population) {
        this.individuals = new ArrayList<>();
        this.individuals.addAll(population.individuals);
    }

    private void researchAlphaIndividual(){
        Iterator<Individual> individualIterator = this.iterator();
        Individual bestIndividual = individualIterator.next();
        Individual currentIndividual;

        while (individualIterator.hasNext()) {
            currentIndividual = individualIterator.next();

            if (currentIndividual.getScore() > bestIndividual.getScore()) {
                bestIndividual = currentIndividual;
            }
        }
        this.alphaIndividual= bestIndividual;
    }
    
    /**
     *
     * @return the best individuals;
     */
    public Individual getAlphaIndividual() {

       this.researchAlphaIndividual();
       return this.alphaIndividual;
    }

    /**
     * @deprecated .
     * @return
     */
    public String xmlSerialisation() {
        String serialisedPopulation = "";
        //TODO
        return serialisedPopulation;
    }

    /**
     * sort the individuals from the best to the worst.
     */
    public void sort() {
        Collections.sort(individuals, new IndividualComparator());
    }

    /**
     *
     * @return
     */
    public int size() {
        return this.individuals.size();
    }
    
    /**
     *
     * @return
     */
    @Override
    public Population clone() {
        
        Population pop = new Population(this.observableVolume);
        pop.addAll(this.individuals);
        
        return pop;
    }
    
    /**
     * remove the individual at the index 'index' in the individuals.
     * @param index
     * @return
     */
    public Individual remove(int index){
        return this.individuals.remove(index);
    }
    
    /**
     *
     * @return true if the popualtion does not contain any individual. False otherwise.
     */
    public boolean isEmpty(){
        return this.individuals.isEmpty();
    }
}
