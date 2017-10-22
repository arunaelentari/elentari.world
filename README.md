# elentari.world

Repo for elentari.world website.

In case you want to deploy your changes in prod, run:

```
bash compileprod && scp elentari.world core@elentari.world:
```
To connect to the prod machine, type:

```
ssh core@elentari.world
```

Every time a server is built, a public-private key pair is generated.
When the server is destroyed, the key pair is destroyed as well.  When
a client wants to connect to the new server, the new key pair is not
recognized, so you need to say yes when connecting to it.

To remove the old host ssh key after rebuilding the server, type:

```
ssh-keygen -R elentari.world
```