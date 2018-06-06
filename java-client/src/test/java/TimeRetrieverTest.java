package se.liu.ida.hello;

import static org.junit.Assert.*;
import junit.framework.TestCase;
import org.joda.time.LocalTime;

public class TimeRetrieverTest extends TestCase {

    public TimeRetrieverTest(String name) {
        super(name);
    }

    public void testTimeRetriever() throws Exception {
        LocalTime before = new LocalTime();
        String testString = "TODO list: " + before.toString();
        String actualString = TimeRetriever.getTimeMessage();
        assertTrue(testString.compareTo(actualString) <= 0);
    }
}