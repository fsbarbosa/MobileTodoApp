import React, { useState, useEffect } from 'react';
import { View, Text, TextInput, StyleSheet, Button, Alert } from 'react-native';
import axios from 'axios';
import { useForm, Controller } from 'react-hook-form';

const EditTaskScreen = ({ route, navigation }) => {
  const { taskId } = route.params;
  const { control, handleSubmit, setValue } = useForm();

  useEffect(() => {
    const fetchTaskDetails = async () => {
      const cachedDetails = localStorage.getItem(taskId);
      if (cachedDetails) {
        const { title, description } = JSON.parse(cachedDetails);
        setValue('title', title);
        setValue('description', description);
        return;
      }
      try {
        const response = await axios.get(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`);
        const { title, description } = response.data;
        
        localStorage.setItem(taskId, JSON.stringify({ title, description }));
        setValue('title', title);
        setValue('description', description);
      } catch (error) {
        console.error("Error fetching task details:", error);
        Alert.alert("Error", "An error occurred while fetching the task details.");
      }
    };

    fetchTaskDetails();
  }, [taskId, setValue]);

  const onSubmit = async (data) => {
    try {
      await axios.put(`${process.env.REACT_APP_BACKEND_URL}/tasks/${taskId}`, data);
      localStorage.setItem(taskId, JSON.stringify(data));
      Alert.alert("Success", "Task updated successfully");
      navigation.goBack();
    } catch (error) {
      console.error("Error updating task:", error);
      Alert.alert("Error", "An error occurred while updating the task.");
    }
  };

  return (
  );
};

export default EditTaskScreen;