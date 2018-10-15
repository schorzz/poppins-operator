# Poppins-Operator

A Poppins is a operator for kubernetes, which enables kubernetes to give namespaces an expiry date. After the namespace is expired,
the poppins-operator will delete deployments and pods. The Operator offers CRUD-operations on poppins ressource definitions (RDs).
Its the goal of this operator to "clean-up" namespaces after they expire. A namespace is expired, if the expiry date is exceeded. This
expiry date is defined in the poppins RD.

Every Namespace, that can expire and needs to be deleted automatically, should have a poppins RD.

This operator offers the possibility to create and update poppinses (poppins RDs). Deleting a poppins is not provided because it
shouldn't be able to delete poppinses via ReST.

## Docs
Docs can be found ->[here](https://schorzz.github.io/poppins-operator/)<-
