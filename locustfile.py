from locust import FastHttpUser, task
import random
import socket

class HelloWorldUser(FastHttpUser):
    def get_base_url(self):
        """Select a host from the list randomly for each request."""
        rand_dst_port = round(random.randint(80, 97))
        return f"http://34.133.88.127:{rand_dst_port}"
    
    @task
    def hello_world(self):
        randuser = round(random.randint(1, 52000))
        self.client.get(f"/todos/{randuser}", base_url=self.get_base_url())
        