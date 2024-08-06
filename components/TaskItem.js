import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity, Alert, TouchableWithoutFeedback } from 'react-native';
import Icon from 'react-native-vector-icons/FontAwesome';

const TaskItem = ({ task, onDeleteTask, onEditTask }) => {
  const handleDelete = async (taskId) => {
    try {
      await onDeleteTask(taskId);
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'An error occurred while deleting the task.');
    }
  };

  const handleEdit = async (task) => {
    try {
      await onEditTask(task);
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'An error occurred while editing the task.');
    }
  };

  const toggleCompletion = async (task) => {
    const updatedTask = { ...task, isCompleted: !task.isCompleted };
    try {
      await onEditTask(updatedTask);
    } catch (error) {
      console.error(error);
      Alert.alert('Error', 'An error occurred while updating the task status.');
    }
  };

  return (
    <View style={styles.taskItemContainer}>
      <TouchableWithoutFeedback onPress={() => toggleCompletion(task)}>
        <Text style={[styles.taskText, task.isCompleted ? styles.taskCompleted : null]}>{task.text}</Text>
      </TouchableWithoutFeedback>
      <TouchableOpacity onPress={() => handleEdit(task)}>
        <Icon name="edit" size={24} color="blue" />
      </TouchableOpacity>
      <TouchableOpacity onPress={() => handleDelete(task.id)}>
        <Icon name="trash" size={24} color="red" />
      </TouchableOpacity>
    </View>
  );
};

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
  taskCompleted: {
    textDecorationLine: 'line-through',
    color: 'grey',
  },
});

export default TaskItem;