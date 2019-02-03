# Tutorials

https://www.youtube.com/watch?v=0CFWeVgzsqY

## Install Maven and dependencies

```sh
brew install maven
# java sdk
brew install gradle
brew install postgresql
```

## Setup project

### pom-root

https://mvnrepository.com/artifact/org.codehaus.mojo.archetypes/pom-root

```sh
mvn archetype:generate -B -DgroupId=com.toblu933 -DartifactId=tdde06-maven -DarchetypeGroupId=org.codehaus.mojo.archetypes  -DarchetypeArtifactId=pom-root -DinteractiveMode=false
```

### client

```sh
mvn archetype:generate -B -DgroupId=com.toblu933 -DartifactId=java-client -DarchetypeGroupId=org.apache.maven.archetypes

# create maven resources folder
mkdir -p java-client/main/resources
```

### server

https://github.com/raydac/mvn-golang

```sh
mvn archetype:generate -B -DarchetypeGroupId=com.igormaznitsa -DarchetypeArtifactId=mvn-golang-hello -DarchetypeVersion=2.2.0 -DgroupId=com.toblu933 -DartifactId=go-server -Dversion=1.0-SNAPSHOT
# inside ./go-server/main/src/
# remove App.java and AppTest.java and clone the ci-sample-project from git
rm App.java AppTest.java
git clone git@gitlab.ida.liu.se:large-scale-dev/ci-sample-project.git .
rm -rf .git # remove git
# add database schema

initdb /usr/local/var/postgres
psql -f schema.sql postgres
pg_ctl -D /usr/local/var/postgres -l logfile start
psql postgres
CREATE USER postgres WITH PASSWORD 'postgress3cre7';
CREATE DATABASE postgres;
ALTER DATABASE postgres OWNER TO postgres;
```

### Higher grade

add log4j.properties to `java-client/src/main/resources

add log4j to java dependencies

```xml
<!-- https://mvnrepository.com/artifact/log4j/log4j -->
<dependency>
    <groupId>log4j</groupId>
    <artifactId>log4j</artifactId>
    <version>1.2.17</version>
</dependency>
```

## Maven commands

```
mvn verify
mvn compile
```

```
# Plugins
https://github.com/raydac/mvn-golang
```
