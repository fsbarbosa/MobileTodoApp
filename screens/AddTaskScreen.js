import React, { useState } from 'react';
import { View, TextInput, Button, StyleSheet, Alert } from 'react-native';
import axios from 'axios';
import { API_URL } from '@env';

const AddTaskScreen = () => {
  const [task, setTask] = useState('');

  const handleAddTask = async () => {
    try {
      if (task.trim() === '') {
        Alert.alert('Validation', 'Please enter a task');
        return;
      }

      const taskData = {
        title: task,
        completed: false
      };

      const response = await axios.post(`${API_URL}/tasks`, taskData);

      if (response.status === 201) {
        setTask('');
        Alert.alert('Success', 'Task added successfully');
      } else {
        Alert.alert('Error', 'Failed to add the task');
      }
    } catch (error) {
      Alert.alert('Error', 'Could not connect to the server');
    }
  };

  return (
    <View style={styles.container}>
      <TextInput
        style={styles.input}
        placeholder="Enter your task here..."
        value={task}
        onChangeText={setTask}
      />
      <Button title="Add Task" onPress={handleAddAddTask} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  input: {
    width: '80%',
    borderBottomWidth: 1,
    borderBottomColor: 'gray',
    margin: 20,
    fontSize: 18,
  },
});

export default AddTaskScreen;