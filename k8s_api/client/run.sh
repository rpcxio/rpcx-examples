https://m0sh1x2.com/posts/optimizing-local-go-kubernetes-deployments/

kubectl create clusterrolebinding pod-reader \
  --clusterrole=pod-reader  \
  --serviceaccount=default:default