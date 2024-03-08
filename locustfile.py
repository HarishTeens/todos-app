from locust import HttpUser, task
import random

class HelloWorldUser(HttpUser):
    @task
    def hello_world(self):
        randuser = round(random.randint(1, 42000))
        self.client.get(f"/todos/{randuser}")