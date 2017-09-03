int getNodeNum(b *root) {
	if(root == null){
		return 0;
	}
	return getNodeNum(root->left) + getNodeNum(root->right) + 1;
}
