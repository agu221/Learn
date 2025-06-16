import numpy as np
from collections import Counter
class TreeNode():
    """
    Tree Node Class
    Contains the individual nodes of the decision tree
    
    Args: 
        data (list): The data passed into the node
        feature_idx (int): Gives the index of the feature in the data set
        feature_val (float): Gives the median value of the feature
        prediciton_probs (List(float)): Gives the prediction probabilities
        information_gain (float): The difference in node entropy and the split entropy 
    """
    def __init__(self, data= None, feature_idx = None, feature_val = None, prediction_probs = None,information_gain =None) -> None:
        self.data = data
        self.feature_idx = feature_idx
        self.feature_val = feature_val
        self.prediction_probs = prediction_probs
        self.information_gain = information_gain
        self.left = None
        self.right = None
        
class DecisionTree():
    """
    Decision Tree Classifer
    Training: Use "train" function with train set features and laebls
    Predicting: Use "predict function with test set features
    
    Args:
        max_depth (int): The maximum number of layers in the tree
        min_samples_leaf (int): The minimum number of the samples per leaf
        min_information_gain (int): The minimum amount of infomration gain to be achived by each split"""
    
    def __init__(self, max_depth: int = 4, min_samples_leaf: int =1, min_information_gain: int = 0.0) -> None:
        self.max_depth = max_depth
        self.min_samples_leaf = min_samples_leaf
        self.min_information_gain = min_information_gain

    def entropy(self,class_of_probabilities:list=None) -> float:
        """
        Given a list of class probabilities, calculate the entropy for that class.

        Args:
            class_of_probabilities (list): A list containing the probabilities of each category for that class
        
        Returns:
            entropy (float): The value of the entropy
        """
        return sum(-p*np.log2(p) for p in class_of_probabilities if p>0)
    
    def find_best_split(self, data: np.array=None)->tuple:
        """
        Given a dataset, the best split of data

        Args:
            data (numpy array): Data set 

        Returns
            splits (tuple): A tuple with g1_min, g2_min, min_entropy_feature_idx, min_entropy_feature_val, min_part_entropy
        """

        min_part_entropy = 1e6
        min_entropy_feautre_idx = None
        min_entropy_feature_val = None
        
        for idx in range(data.shape[1]-1):
            feature_val = np.median(data[:,idx])
            g1,g2 = self.split(data,idx,feature_val)
            part_entropy = self.partition_entropy([g1[:,-1],g2[:,-1]])
            if part_entropy < min_part_entropy:
                min_part_entropy = part_entropy
                min_entropy_feature_val = feature_val
                min_entropy_feautre_idx = idx
                g1_min, g2_min = g1, g2
        
        return g1_min, g2_min, min_entropy_feautre_idx, min_entropy_feature_val, min_part_entropy
    
    def create_tree(self, data: np.array, current_depth:int)->TreeNode:
        """
        A recurrsive function to build a tree based on given stopping criteria
        Args:
            data (numpy array): A dataset
            current_depth (int): The current depth of the recurrsion
        """
        if current_depth >= self.max_depth:
            return None
        split_1_data, split_2_data, split_feature_idx, split_feature_val, split_entropy = self.find_best_split(data)

        label_probabilities = self.find_label_probs(data)

        node_entropy = self.entropy(label_probabilities)
        information_gain = node_entropy - split_entropy

        node = TreeNode(data,split_feature_idx,split_feature_val,label_probabilities,information_gain)

        if self.min_samples_leaf > split_1_data.shape[0] or self.min_samples_leaf > split_2_data.shape[0]:
            return node
        elif information_gain < self.min_information_gain:
            return node
        
        current_depth += 1
        node.right = self.create_tree(split_1_data,current_depth)
        node.left = self.create_tree(split_2_data,current_depth)

        return node
    
    def split(self, data:np.array,feature_idx:int,feature_val:float)->tuple:
        """
        Function to split data given the feature index and feature value.
        
        Args:
            data (numpy array): The data set
            feature_idx (int): The index of where to split the feature
            feature_val (float): The value to use to split the groups
        """
        mask_below_threshold = data[:,feature_idx] < feature_val
        group1 = data[mask_below_threshold]
        group2 = data[~mask_below_threshold]

        return group1 , group2
    
    def find_label_probs(self,data:np.array)->np.array:
        """
        Function to find the probability of each label of a data set.
        Args:
            data (numpy array): The data set
        Returns
            label_probabilities (numpy array): The probability of each label
        """
        labels_as_integers = data[:,-1].astype(int)
        total_labels = len(labels_as_integers)
        label_probabilities = np.zeros(len(self.labels_in_train),dtype=float)

        for i, label in enumerate(self.labels_in_train):
            label_index = np.where(labels_as_integers ==i)[0]
            if len(label_index>0):
                label_probabilities[i] = len(label_index)/total_labels
        
        return label_probabilities
    
    def partition_entropy(self, subsets:list) ->float:
        """
        Function to find the entropy of a partition which is the entropy
        """
        total_count = sum([len(subset) for subset in subsets])
        return sum(self.data_entropy(subset)*(len(subset)/total_count) for subset in subsets)
    
    def data_entropy(self, labels: list)->float:
        return self.entropy(self.class_probabilities(labels))
    
    def class_probabilities(self, labels:list)->list:
        total_count = len(labels)
        return [label_count/total_count for label_count in Counter(labels).values()]


        

test = np.array([[0.5,0.5,0.5,0.5]])


dt = DecisionTree()

print(dt.partition_entropy(test))