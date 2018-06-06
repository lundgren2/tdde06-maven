package se.liu.ida.hello;

import org.joda.time.LocalTime;
import org.apache.log4j.Logger;

public class TimeRetriever {
    public static String getTimeMessage() {
        Logger logger = Logger.getLogger(HelloMaven.class);
        LocalTime localTime = new LocalTime();

        logger.info("Requested time string: " + localTime.toString());
        return "TODO list: " + localTime.toString();
    }
}