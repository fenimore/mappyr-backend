#from nose.tools import assert_true
import requests

token_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiMSIsImV4cCI6MTQ3Njk3MTQ2NywiaXNzIjoibG9jYWxob3N0OjgwODAifQ.wv1UgclK5uKUYYZFpnx4DjLcHlinirTfzL0nhmrJ7gc"

token_8 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiOCIsImV4cCI6MTQ3Njk3MzEwNCwiaXNzIjoibG9jYWxob3N0OjgwODAifQ.NnxhEH7ETklp8H_hkNNxFaJpHt0s4TdE2gWVRpSE39Q"

def signup():
    """Create a new user"""
    r = requests.post('http://localhost:8080/signup', json ={"username":"Wrinkle",
                                                             "password":"Woootwooot",
                                                             "email":"Noootme@gmail.com"})
    print(r.status_code)
    print(r.json())


#def login():

def comment():
    """Post a comment with a token in the headers"""
    r = requests.post('http://localhost:8080/new', json={"title":"SUsh1","description":"THIS Rox0rs", "latitude":45, "longitude":88}, headers={"Authentication":token_8})

    print(r.status_code)
    print(r.json())
