package se.liu.ida.hello;

import static org.junit.Assert.*;
import junit.framework.TestCase;

import org.apache.http.client.methods.HttpGet;
import org.apache.http.client.HttpClient;
import org.apache.http.impl.client.DefaultHttpClient;
import org.apache.http.HttpResponse;

public class StatusIT extends TestCase {

    private static String serverURL = "http://localhost:8080/list";

    public StatusIT(String name) {
        super(name);
    }

    public void testStatusCode() throws Exception {
        DefaultHttpClient httpClient = new DefaultHttpClient();
        HttpGet getRequest = new HttpGet(serverURL);
        getRequest.addHeader("accept", "application/json");
        HttpResponse response = httpClient.execute(getRequest);

        assertTrue(response.getStatusLine().getStatusCode() == 200);
    }
}