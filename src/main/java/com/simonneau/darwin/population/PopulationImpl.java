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
package com.simonneau.darwin.population;

import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.Iterator;

/**
 *
 * @author simonneau
 * @param <T>
 */
public class PopulationImpl<T extends Individual> implements Population<T>{

    private ArrayList<T> individuals;
    private int populationSize;
    
    public PopulationImpl() {
        this(1);
    }

    public PopulationImpl(int populationSize) {
        this.individuals = new ArrayList<>();
        this.populationSize = populationSize;
    }

    @Override
    public void setPopulationSize(int populationSize) {
        if (populationSize < 1) {
            throw new IllegalArgumentException("population size must be > 0");
        }
        this.populationSize = populationSize;
    }
    
    @Override
    public T getAlphaIndividual() {
       return Collections.max(this.individuals);
    }

    @Override
    public int getPopulationSize() {
        return this.populationSize;
    }

    @Override
    public void sort() {
        Collections.sort(this.individuals, new IndividualComparator());
    }

    @Override
    public int size() {
        return this.individuals.size();
    }

    @Override
    public boolean isEmpty() {
        return this.individuals.isEmpty();
    }

    @Override
    public boolean contains(Object o) {
        return this.individuals.contains(o);
    }

    @Override
    public Iterator<T> iterator() {
        return this.individuals.iterator();
    }

    @Override
    public Object[] toArray() {
        return this.individuals.toArray();
    }

    @Override
    public <T> T[] toArray(T[] a) {
        return this.individuals.toArray(a);
    }

    @Override
    public boolean add(T e) {
        return this.individuals.add(e);
    }

    @Override
    public boolean remove(Object o) {
        return this.individuals.remove(o);
    }

    @Override
    public boolean containsAll(Collection<?> c) {
        return this.individuals.containsAll(c);
    }

    @Override
    public boolean addAll(Collection<? extends T> c) {
        return this.individuals.addAll(c);
    }

    @Override
    public boolean removeAll(Collection<?> c) {
        return this.individuals.removeAll(c);
    }

    @Override
    public boolean retainAll(Collection<?> c) {
        return this.individuals.retainAll(c);
    }

    @Override
    public void clear() {
        this.individuals.clear();
    }

    
    @Override
    public Population<T> clone(){
        PopulationImpl<T> clone = new PopulationImpl<>(this.populationSize);
        clone.individuals = (ArrayList<T>) this.individuals.clone();
        return clone;
    }
    
}
