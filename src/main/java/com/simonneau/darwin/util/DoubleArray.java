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
package com.simonneau.darwin.util;

/**
 * Extrait de la librairie qui ne contenait que cette méthode utile.
 *
 * @author Yann RICHET
 */
public class DoubleArray {
    /**
     * Generates an array of successive values from <em>begin</em> to <em>end</em> with step
     * size <em>pitch</em>.
     * @param begin First value in sequence
     * @param pitch Step size of sequence
     * @param end Last value of sequence
     * @return Array of successive values
     */
    public static double[] increment(double begin, double pitch, double end) {
        double[] array = new double[(int) (((end - begin) / pitch)+1)];
        for (int i = 0; i < array.length; i++) {
            array[i] = begin + i * pitch;
        }
        return array;
    }
}
