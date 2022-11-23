#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>

const char* ssid = "SSID";
const char* password = "PASSWORD";

void setup() {
  Serial.begin(115200);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.println("WiFi connected");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());
}

void loop() {
  HTTPClient http;
  http.begin("http://https://arduino-beskend.herokuapp.com/on");
    //post request {"onLamp": true, "onMator": true}
    http.addHeader("Content-Type", "application/json");
    int httpCode = http.POST("{\"onLamp\": true, \"onMator\": true}");
    Serial.println(httpCode);
    http.end();
    delay(5000);
}

