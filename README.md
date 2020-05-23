# kubernetes-dependon-operator
Adding pod dependency to kubernetes.For example if a web server needs to a database it should watch and if database shut , down also web service must shut down because of not wasting resources.<br><br>
For installing you should first build docker image<br>
docker build -t anilkuscu95/dependon .<br>
docker push anilkuscu95/dependon<br><br>
After push apply the yaml file on kubernetes.<br><br><br>
kubectl apply -f dependon.yaml<br><br>
You should edit dependon.yaml file as your will.It includes a example dependon crd on the last section.<br>
