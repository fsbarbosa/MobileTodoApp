import React, { useState } from 'react';
import { View, Text, StyleSheet, FlatList, TouchableOpacity } from 'react-native';
import { useNavigation } from '@react-navigation/native';

const SAMPLE_TASKS = [
  { id: '1', title: 'Complete React Native App', description: 'Finish the code for the Todo List app' },
  { id: '2', title: 'Grocery Shopping', description: 'Buy milk, eggs, and bread for the week' },
];

const HomeScreen = () => {
  const navigation = useNavigation();
  const [tasks, setTasks] = useState(SAMPLE_TASKS);

  const TaskItem = ({ item }) => (
    <TouchableOpacity
      style={styles.taskContainer}
      onPress={() => navigation.navigate('EditTask', { taskId: item.id })}
    >
      <Text style={styles.taskTitle}>{item.title}</Text>
    </TouchableOpacity>
  );

  return (
    <View style={styles.container}>
      <FlatList
        data={tasks}
        renderItem={TaskItem}
        keyExtractor={item => item.id}
      />
      <TouchableOpacity
        style={styles.addTaskButton}
        onPress={() => navigation.navigate('AddTask')}
      >
        <Text style={styles.addTaskButtonText}>+</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  taskContainer: {
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#ddd',
  },
  taskTitle: {
    fontSize: 18,
  },
  addTaskButton: {
    position: 'absolute',
    right: 20,
    bottom: 20,
    width: 60,
    height: 60,
    borderRadius: 30,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: 'blue',
  },
  addTaskButtonText: {
    color: '#fff',
    fontSize: 30,
  },
});

export default HomeScreen;