# DB_GAUVIN_P01
Shop DB in Golang with hot reload system

## Requirements

If you use docker you will only need:
* Docker;
* Docker-Compose;

Refer to [Docker-Setup](#docker-setup) to install with docker.

If not, to run this project, you will need to install the following dependencies on your system:

- [go](https://golang.org/doc/install)

## Docker-Setup

* Run with : ```docker-compose up --build```

After docker run, following command :

* Import db : ``` docker exec -i db_gauvin_p01_db_1 sh -c 'exec mysql -ugobdd -p"gobdd" image_gobdd' < ./docker/data/database.sql```

!! On linux, if you have a permission denied error on mysql_data, run : ``` sudo chown -R <user>:<user> ./<folder> ```

## API Documentation

You can acces the API documentation here : https://documenter.getpostman.com/view/12952136/TVRoYRr2

## Linter

We use go linter [gofmt](https://blog.golang.org/gofmt) to automatically formats the source code.

## Contributors

<table>
  <tr>
    <td align="center">
    <a href="https://github.com/jasongauvin">
      <img src="https://avatars1.githubusercontent.com/u/41618366?s=400&u=b970ed03cbb921ce1312ef86b39093e4fa0be7e3&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jason Gauvin</b></sub>
    </a>
    </td>
    <!-- <td align="center">
    <a href="https://github.com/JackMaarek/">
      <img src="https://avatars3.githubusercontent.com/u/28316928?s=400&u=3cdfb5b0683245ad333a39cfca3a5251f3829824&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jacques Maarek</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/SteakBarbare">
      <img src="https://avatars2.githubusercontent.com/u/25483831?s=400&u=5316e2018489cb088c6120940df7e0b5d8d0f374&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Corto Dufour</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/edwinvautier">
      <img src="https://avatars3.githubusercontent.com/u/35581502?s=460&u=d9096f90151f35552d9adcd57bacaee366f0aaef&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Edwin Vautier</b></sub>
    </a>
    </td> -->
  </tr>
</table>
