import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, StyleSheet, Button } from 'react-native';
import axios from 'axios';
import { useForm, Controller } from 'react-hook-form';

const EditTaskScreen = ({ route, navigation }) => {
  const { taskId } = route.params;
  const { control, handleSubmit, setValue } = useForm();

  useEffect(() => {
    const fetchTaskDetails = async () => {
      const response = await axios.get(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`);
      const { title, description } = response.data;
      setValue('title', title);
      setValue('description', description);
    };
    
    fetchTaskDetails();
  }, [taskId, setValue]);

  const onSubmit = async (data) => {
    try {
      await axios.put(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`, data);
      navigation.goBack();
    } catch (error) {
      console.error("Error updating task:", error);
    }
  };

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Edit Task</Text>
      <Controller
        control={control}
        render={({ field: { onChange, onBlur, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            placeholder="Title"
          />
        )}
        name="title"
        rules={{ required: true }}
        defaultValue=""
      />
      <Controller
        control={control}
        render={({ field: { onChange, onBlur, value } }) => (
          <TextInput
            style={styles.input}
            onBlur={onBlur}
            onChangeText={onChange}
            value={value}
            placeholder="Description"
            multiline={true}
          />
        )}
        name="description"
        rules={{ required: true }}
        defaultValue=""
      />
      <Button title="Submit" onPress={handleSubmit(onSubmit)} />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: 'center',
    justifyContent: 'center',
    padding: 20,
  },
  title: {
    fontSize: 24,
    fontWeight: 'bold',
    marginBottom: 20,
  },
  input: {
    width: '100%',
    borderWidth: 1,
    borderColor: 'gray',
    borderRadius: 5,
    padding: 10,
    marginBottom: 20,
    fontSize: 18,
  },
});

export default EditTaskScreen;