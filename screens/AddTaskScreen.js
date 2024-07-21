import React, { useState } from 'react';
import { View, TextInput, Button, StyleSheet, Alert, Text, ScrollView, ActivityIndicator } from 'react-native';
import axios from 'axios';
import { API_URL } from '@env';

const AddTaskScreen = () => {
  const [task, setTask] = useState('');
  const [tasks, setTasks] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleAddTask = async () => {
    try {
      if (task.trim() === '') {
        Alert.alert('Validation', 'Please enter a task');
        return;
      }

      setIsLoading(true);

      const taskData = {
        title: task,
        completed: false
      };

      const response = await axios.post(`${API_URL}/tasks`, taskData);

      if (response.status === 201) {
        setTasks(currentTasks => [...currentTasks, { ...taskData, id: response.data.id }]);
        setTask('');
        Alert.alert('Success', 'Task added successfully');
      } else {
        Alert.alert('Error', 'Failed to add the task');
      }
    } catch (error) {
      Alert.alert('Error', 'Could not connect to the server');
    } finally {
      setIsLoading(false);
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
      { isLoading ? (
        <ActivityIndicator size="large" />
      ) : (
        <Button title="Add Task" onPress={handleAddType} />
      )} 
      <Text style={styles.tasksHeader}>Tasks Added:</Text>
      <ScrollView>
        {tasks.map((item, index) => (
          <View key={item.id || index} style={styles.taskItem}>
            <Text style={styles.taskText}>{item.title}</Text>
          </View>
        ))}
      </ScrollView>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    paddingTop: 30,
  },
  input: {
    width: '80%',
    borderBottomWidth: 1,
    borderBottomColor: 'gray',
    margin: 20,
    fontSize: 18,
  },
  tasksHeader: {
    fontSize: 20,
    marginTop: 20,
    marginBottom: 10,
  },
  taskItem: {
    backgroundColor: '#f9c2ff',
    padding: 20,
    marginVertical: 4,
    marginHorizontal: 16,
    borderRadius: 10,
  },
  taskText: {
    fontSize: 18,
  },
});

export default AddTaskScreen;