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

import java.util.GregorianCalendar;

/**
 *
 * @author nono
 */
public class Chronometer {

    private GregorianCalendar date;
    private long currentSessionTimeCount = 0;
    private long previousSessionTimeCount = 0;
    private boolean stoped = true;

    
    /**
     *
     */
    public void reset(){
        this.stop();
        this.setTime(0);
    }    
    /**
     * restart 'this' chronometer.
     */
    public void restart() {
        this.setTime(0);
        this.start();
    }

    /**
     * start 'this' chronometer.
     */
    public void start() {
        this.date = new GregorianCalendar();
        this.stoped = false;
    }

    /**
     * stop 'this' chronometer.
     */
    public void stop() {
        if (!this.stoped) {
            this.previousSessionTimeCount += new GregorianCalendar().getTimeInMillis()- this.date.getTimeInMillis();
            this.stoped = true;
        }
    }

    /**
     * set 'this' current session time.
     * @param time
     */
    private void setTime(int time) {
        this.currentSessionTimeCount = time;
        this.previousSessionTimeCount = 0;
    }

    /**
     *
     * @return 'this' current session time.
     */
    public long getTime() {
        if (!this.stoped) {
            this.currentSessionTimeCount = new GregorianCalendar().getTimeInMillis()- this.date.getTimeInMillis();
        }
        return this.currentSessionTimeCount + this.previousSessionTimeCount;
    }
}
