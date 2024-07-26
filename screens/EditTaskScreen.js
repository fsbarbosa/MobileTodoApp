import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, StyleSheet, Button, Alert } from 'react-native';
import axios from 'axios';
import { useForm, Controller } from 'react-hook-form';

const EditTaskScreen = ({ route, navigation }) => {
  const { taskId } = route.params;
  const { control, handleSubmit, setValue, getValues } = useForm();

  useEffect(() => {
    const fetchTaskDetails = async () => {
      const response = await axios.get(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`);
      const { title, description } = response.data;
      setValue('title', title);
      setValue('description', description);
    };

    fetchTaskDetails();
  }, [taskId, setValue]);

  const validateAndSubmit = async (data) => {
    if (!data.title.trim() || !data.description.trim()) {
      Alert.alert("Validation Error", "Title and Description cannot be empty.");
      return;
    }
    onSubmit(data);
  };

  const onSubmit = async (data) => {
    try {
      await axios.put(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`, data);
      Alert.alert("Success", "Task updated successfully");
      navigation.goBack();
    } catch (error) {
      console.error("Error updating task:", error);
      Alert.alert("Error", "An error occurred while updating the task.");
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
      <Button title="Submit" onPress={handleSubmit(validateAndSubmit)} />
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