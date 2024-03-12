from locust import FastHttpUser, task
import random

class HelloWorldUser(FastHttpUser):
    @task
    def hello_world(self):
        randuser = round(random.randint(1, 52000))
        self.client.get(f"/todos/{randuser}")