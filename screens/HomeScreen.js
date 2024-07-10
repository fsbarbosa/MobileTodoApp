import React, {useState, useEffect} from 'react';
import {View, Text, StyleSheet, FlatList, TouchableOpacity} from 'react-native';
import {useNavigation} from '@react-navigation/native';

const SAMPLE_TASKS = [
  {id: '1', title: 'Task 1', description: 'Description of Task 1'},
  {id: '2', title: 'Task 2', description: 'Description of Task 2'},
];

const HomeScreen = () => {
  const navigation = useNavigation(); 
  const [tasks, setTasks] = useState(SAMPLE_TASKS); 

  useEffect(() => {
  }, []); 

  const renderTask = ({item}) => (
    <TouchableOpacity
      style={styles.taskItem}
      onPress={() => navigation.navigate('EditTask', {taskId: item.id})} 
    >
      <Text style={styles.taskTitle}>{item.title}</Text>
    </TouchableOpacity>
  );

  return (
    <View style={styles.container}>
      <FlatList
        data={tasks} 
        renderItem={renderTask} 
        keyExtractor={item => item.id} 
      />
      <TouchableOpacity
        style={styles.addButton}
        onPress={() => navigation.navigate('AddTask')} 
      >
        <Text style={styles.addButtonText}>+</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  },
  taskItem: {
    padding: 20,
    borderBottomWidth: 1,
    borderBottomColor: '#ddd',
  },
  taskTitle: {
    fontSize: 18,
  },
  addButton: {
    position: 'absolute',
    right: 20,
    bottom: 20,
    width: 60,
    height: 60,
    borderRadius: 30,
    justifyContent: 'center',
    alignItems: 'center',
    backgroundColor: '#blue',
  },
  addButtonText: {
    color: '#fff',
    fontSize: 30,
  },
});

export default HomeScreen;