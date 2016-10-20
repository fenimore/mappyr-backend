#from nose.tools import assert_true
import requests

# Example tokens with id
token_1 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiMSIsImV4cCI6MTQ3Njk3MTQ2NywiaXNzIjoibG9jYWxob3N0OjgwODAifQ.wv1UgclK5uKUYYZFpnx4DjLcHlinirTfzL0nhmrJ7gc"

token_8 = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyLWlkIjoiOCIsImV4cCI6MTQ3Njk3MzEwNCwiaXNzIjoibG9jYWxob3N0OjgwODAifQ.NnxhEH7ETklp8H_hkNNxFaJpHt0s4TdE2gWVRpSE39Q"



def set_up():
    """This sets up a nice mock database with some fancy stuffs, and prints some rezults"""
    user_1 = {"username":"Wrinkle", "password":"Woootwooot", "email":"Noootme@gmail.com"}
    user_2 = {"username":"Simone", "password":"lulz", "email":"simone@gmail.com"}
    user_3 = {"username":"Anon", "password":"test", "email":"anon@gmail.com"}
    u1 = signup(user_1)
    u2 = signup(user_2)
    u3 = signup(user_3)

    print("Test User 2 ", u2)

    simone_token = login({"username":"Simone", "password":"lulz"})
    anon_token = login({"username":"Anon", "password":"test"})
    # The token response is a json/dict, with one field:
    # Authentication : TOKEN

    print("Test Token Anon ", anon_token["Authentication"])

    c_1 = comment({"title":"SUsh1","description":"THIS Rox0rs",
                            "latitude":45, "longitude":88}, simone_token["Authentication"])
    c_2 = comment({"title":"Park Bench","description":"THIS sux0rs",
                            "latitude":43.44444, "longitude":84}, anon_token["Authentication"])

    print("Test Comment 1", c_1)

    downvoted = downvote(c_2["id"], simone_token["Authentication"])
    upvoted = upvote(c_1["id"], anon_token["Authentication"])


    print("\n\nSUCCESS UNTIL HERE\n\n")
    print(user_votes(u1["id"]))
    print(user_votes(u2["id"]))
    print(user_votes(u3["id"]))

##########################################################################################
#USER ACTIONS
##########################################################################################
def signup(j):
    """Create a new user

    returns the user in json
    """
    r = requests.post('http://localhost:8080/signup', json=j)

    return r.json()


def login(j):
    """Login a user with a username and password

    returns the auth token in dict
    """
    r = requests.post('http://localhost:8080/login', json=j)

    return r.json()

def comment(j, token):
    """Post a comment with a token in the headers"""
    r = requests.post('http://localhost:8080/new',
                      json=j,
                      headers={"Authentication":token})

    return r.json()


def upvote(comment_id, token):
    """UpVOTE!"""
    r = requests.get('http://localhost:8080/upvote/'+ str(comment_id), headers={"Authentication":token})
    return r.json()

def downvote(comment_id, token):
    """downVOTE!"""
    r = requests.get('http://localhost:8080/downvote/'+ str(comment_id), headers={"Authentication":token})
    return r.json()

##########################################################################################
#Get Votes or COmments of User
##########################################################################################
def user_votes(id):
    """User votes takes the id, but pass a string"""
    r = requests.get("http://localhost:8080/votes/"+str(id))

    return r.json()

def user_comments(id):
    """User comments takes the id, but pass a string"""
    r = requests.get("http://localhost:8080/comments/"+str(id))

    return r.json()
