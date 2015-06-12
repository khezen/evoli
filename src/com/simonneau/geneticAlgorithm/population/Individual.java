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

/**
 *
 * @param <T> 
 * @author simonneau
 */
public abstract class Individual<T extends Individual> implements Comparable<T> {

    private double score;

    /**
     * Overall properties of 'this' with 's' properties.
     *
     * @param s
     */
    protected abstract void set(Individual s);

    /**
     *
     * @return 'this' score.
     */
    public double getScore() {
        return this.score;
    }

    /**
     * set'this' score.
     * @param score
     */
    public void setScore(double score) {
        this.score = score;
    }

    /**
     *
     * @return
     */
    public abstract String xmlSerialization();

    /**
     * compare this to t.
     * @param t
     * @return 1 if this.getScore() > t.getScore(). 0 if this.getScore() == t.getScore(). -1 otherwise.
     */
    @Override
    public int compareTo(Individual t) {
        
        double thisScore = this.getScore();
        double tScore = t.getScore();

        if (thisScore > tScore) {
            return 1;
        } else if (thisScore < tScore) {
            return -1;
        } else {
            return 0;
        }
    }
}
