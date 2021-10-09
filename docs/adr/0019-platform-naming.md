# 19. Platform naming

Date: 2021-10-09

## Status

Status: Accepted on 2021-10-09

## Context

There are different names and views from different perspectives on the platform, we need to make clear how the naming is used and for what.

## Decision

A user of the platform will use a name for the platform which in most cases is tied to either an environment like the test, prod, etc.. or a business purpose like banking, shop, etc... This name is a purpose name from the user's perspective and is important when communicating to the user but it isn't a name that needs to carry through the rest of Opsteady setup.

A second semi-important name is a subdomain for the platform which will mostly be the same as the user's platform name for easier recognition but can be adjusted for any reason.

A platform runs in an environment (Azure subscription or AWS account) that doesn't really have a name by itself but it should have something recognizable for Opsteady operators and the users. Therefore it makes sense to use the same name as the user does. This name is an identifier and can be changed, it is used in Vault to identify the environment so keep that in mind.

Using the same name for the above makes communication much easier between Opsteady operators and users, it shows a clear purpose and makes things unique. This name is only an identifier and can be changed if needed because for example the purpose of the platform changes. Here are some examples of the names that could be used by users dev-azure, demo-aws, shop-prod, userX-prod, etc...

Besides an identifier, all the resources in the cloud have to have a name as well. It would be okay to use the same name the users use here as well but the drawback is that when you switch between environments the naming of the resources is constantly different. This can be fixed by using a prefix and adding the name as a suffix. However when the intent of the platform changes, as described above, the name of the resources can't change because it means the resources need to be recreated in most cases. Therefore the resource names should be decoupled from the name the users use.
Another option would be to give a name to a resource without being dynamic, without any dynamic prefix or suffix, however limits the option to create another platform into the same environment in the future. To create consistency and tackle this problem we are going to use `plt1` as the name or as a suffix name for resources created in the cloud. This way the naming is always consistent and we can create another platform in the environment by adding an extra number from plt1 to plt2. This is mostly useful in a dev situation where having multiple Azure subscriptions or AWS accounts isn't always desired.

## Consequences

We create flexibility and rigidness with this approach as we are fixating the name of the resources to an arbitrary name plt1.
