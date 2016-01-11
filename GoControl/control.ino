int Delay = 1000;

void setup() {
  Serial.begin(9600);
  pinMode(13, OUTPUT);
}

void loop() {
  if (Serial.available()) {
    Delay = Serial.read();
  }
  // put your main code here, to run repeatedly:
  digitalWrite(13, HIGH);
  delay(Delay);
  digitalWrite(13, LOW);
  delay(Delay);
}
