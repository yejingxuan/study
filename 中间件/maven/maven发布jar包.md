mvn deploy:deploy-file -DgroupId=com.yjx -DartifactId=yjx-demo -Dversion=1.2.3-beta -Dpackaging=jar -Dfile=E:/xx.jar -Durl=http://xxx/repository/maven-releases/ -DrepositoryId=Releases


mvn deploy:deploy-file -DgroupId=com.yjx -DartifactId=yjx-demo -Dversion=dzh-1.1-SNAPSHOT -Dpackaging=jar -Dfile=E:/xxx.jar -Durl=http://xxx/repository/maven-snapshots/ -DrepositoryId=Snapshots