### Start Azure CLI 2.0 from Docker
```
$ docker run -v "$PWD":/root -it azuresdk/azure-cli-python
bash-4.3# cd
```

## Login using the CLI

```
az login -u "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" -p "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX=" --service-principal --tenant "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
```

## Read the environement variables

```
bash-4.3# source 00_environment.sh

The resource group is set to: YOUR-RESOURCE-GROUP
Service plan is set to: YourAppServicePlan.

We are going to deploy in region: westeurope

bash-4.3#
```

### Create the resource group
```
bash-4.3# ./01_create_resource_group.sh
```

### Create the App Service Plan
```
bash-4.3# ./02_create_app_service.sh
```

### Create the Web App
```
bash-4.3# ./03_create_webapp.sh
```