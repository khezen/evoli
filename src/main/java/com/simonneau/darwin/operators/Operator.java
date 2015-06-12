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
public abstract class Operator{

    /**
     *
     */
    protected String label;
    
    /**
     *
     * @param label
     */
    public Operator(String label){
        this.label = label;
    }
    
    /**
     *
     * @return 'this' label.
     */
    public String getLabel() {
        return this.label;
    }
    
    /**
     *
     * @return 'this' label.
     */
    @Override
    public String toString(){
        return this.getLabel();
    }
}