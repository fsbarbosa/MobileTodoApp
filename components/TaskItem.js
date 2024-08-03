import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Alert } from 'react-native';
import Icon from 'react-native-vector-icons/FontAwesome';

const TaskItem = ({ task, onDeleteTask, onEditTask }) => {
  // Pseudo error handler to demonstrate the concept
  const handleDelete = async (taskId) => {
    try {
      await onDeleteTask(taskId); // Assuming onDeleteTask returns a Promise
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'An error occurred while deleting the task.');
    }
  };

  const handleEdit = async (task) => {
    try {
      await onEditTask(task); // Assuming onEditTask returns a Promise
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'An error occurred while editing the task.');
    }
  };

  return (
    <View style={styles.taskItemContainer}>
      <Text style={styles.taskText}>{task.text}</Text>
      <TouchableOpacity onPress={() => handleEdit(task)}>
        <Icon name="edit" size={24} color="blue" />
      </TouchableOpacity>
      <TouchableOpacity onPress={() => handleDelete(task.id)}>
        <Icon name="trash" size={24} color="red" />
      </TouchableOpacity>
    </View>
  );
};

// Styles remain unchanged
const styles = StyleSheet.create({
  taskItemContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    justifyContent: 'space-between',
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#ccc',
  },
  taskText: {
    flex: 1,
  },
});

export default TaskItem;