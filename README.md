### WASAPhoto

Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can
upload your photos directly from your PC, and they will be visible to everyone following you.

## Functional design specifications

Each user will be presented with a stream of photos (images) in reverse chronological order, with
information about when each photo was uploaded (date and time) and how many likes and comments
it has. The stream is composed by photos from “following” (other users that the user follows). Users
can place (and later remove) a “like” to photos from other users. Also, users can add comments to any
image (even those uploaded by themself). Only authors can remove their comments.
Users can ban other users. If user Alice bans user Eve, Eve won’t be able to see any information about
Alice. Alice can decide to remove the ban at any moment.
Users will have their profiles. The personal profile page for the user shows: the user’s photos (in reverse
chronological order), how many photos have been uploaded, and the user’s followers and following.
Users can change their usernames, upload photos, remove photos, and follow/unfollow other users.
Removal of an image will also remove likes and comments.
A user can search other user profiles via username.
A user can log in just by specifying the username. 

## How to run  (in development mode)

For backend run the following: go run ./cmd/webapi/

For frontend run in new terminal: ./open-npm.sh and then inside the container npm run dev

## How to run (with containers)

build the images:
docker build -t wasa-photos-backend:latest -f Dockerfile.backend .
docker build -t wasa-photos-frontend:latest -f Dockerfile.frontend .

run the container images:
docker run -it --rm -p 3000:3000 wasa-photos-backend:latest
docker run -it --rm -p 8081:80 wasa-photos-frontend:latest