func (r *DiskPressureHandlerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    log := r.Log.WithValues("diskpressurehandler", req.NamespacedName)

    // Fetch the DiskPressureHandler instance
    var dpHandler opsv1alpha1.DiskPressureHandler
    if err := r.Get(ctx, req.NamespacedName, &dpHandler); err != nil {
        if errors.IsNotFound(err) {
            return ctrl.Result{}, nil
        }
        return ctrl.Result{}, err
    }

    // List all nodes in the cluster
    nodes := &corev1.NodeList{}
    if err := r.List(ctx, nodes); err != nil {
        return ctrl.Result{}, err
    }

    for _, node := range nodes.Items {
        for _, condition := range node.Status.Conditions {
            if condition.Type == corev1.NodeDiskPressure && condition.Status == corev1.ConditionTrue {
                log.Info("Node has disk pressure", "node", node.Name)

                // Trigger Docker prune command via an exec or job
                if err := r.cleanDockerFiles(ctx, &node); err != nil {
                    log.Error(err, "Failed to clean Docker files", "node", node.Name)
                }
            }
        }
    }

    return ctrl.Result{}, nil
}

func (r *DiskPressureHandlerReconciler) cleanDockerFiles(ctx context.Context, node *corev1.Node) error {
    // Use SSH or Kubernetes Job to run the Docker prune command
    cmd := []string{"docker", "system", "prune", "-af"}

    // Execute command on node (example using an SSH library or Kubernetes Job)
    // Example using exec in Kubernetes (simplified for illustration)
    // Your implementation might vary
    err := execCommandOnNode(node, cmd)
    if err != nil {
        return err
    }

    return nil
}
