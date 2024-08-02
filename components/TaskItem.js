import React from 'react';
import { View, Text, StyleSheet, TouchableOpacity } from 'react-native';
import Icon from 'react-native-vector-icons/FontAwesome';
const TaskItem = ({ task, onDeleteTask, onEditTask }) => {
  return (
    <View style={styles.taskItemContainer}>
      <Text style={styles.taskText}>{task.text}</Text>
      <TouchableOpacity onPress={() => onEditTask(task)}>
        <Icon name="edit" size={24} color="blue" />
      </TouchableOpacity>
      <TouchableOpacity onPress={() => onDeleteTask(task.id)}>
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
});
export default TaskItem;