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
package com.simonneau.darwin.problem;

/**
 *
 * @author simonneau
 */
public class StopCriteria{
     private int maxStepCount;
     private int timeout;
     private double evolutionCoeff;
     
     
     /**
     *
     */
    public StopCriteria(){
         this(0,0,0);
     }
     
     /**
     *
     * @param maxStepCount
     * @param timeout
     * @param mineEvolutionCoeff
     */
    public StopCriteria(int maxStepCount, int timeout, double mineEvolutionCoeff ){
         this.maxStepCount = maxStepCount;
         this.timeout = timeout;
         this.evolutionCoeff = mineEvolutionCoeff;         
     }

    /**
     *
     * @return
     */
    public int getMaxStepCount() {
        return maxStepCount;
    }
    
    /**
     *
     * @return
     */
    public int getTimeout() {
        return timeout;
    }

    /**
     *
     * @return
     */
    public double getEvolutionCriterion() {
        return evolutionCoeff;
    }

    /**
     *
     * @param maxStepCount
     */
    public void setMaxStepCount(int maxStepCount) {
        this.maxStepCount = maxStepCount;
    }

    /**
     *
     * @param timeout
     */
    public void setTimeout(int timeout) {
        this.timeout = timeout;
    }

    /**
     *
     * @param minEvolutionCoeff
     */
    public void setMinEvolutionCriterion(double minEvolutionCoeff) {
        this.evolutionCoeff = minEvolutionCoeff;
    }
     
    /**
     *
     * @param stepCount
     * @param time
     * @param evolutionCoeff
     * @return
     */
    public boolean areReached(int stepCount, long time, double evolutionCoeff){
        boolean areReached = false;
        
        if(this.maxStepCount != 0 && stepCount >= this.maxStepCount){
            areReached = true;
            
        }else if(this.timeout != 0 && time >= this.timeout){
            areReached = true;
            
        }else if(this.evolutionCoeff != 0 && evolutionCoeff <= this.evolutionCoeff){
            areReached = true;
        }
        
        return areReached;
    }
     
}
