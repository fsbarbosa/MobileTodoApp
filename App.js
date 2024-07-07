import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import HomeScreen from './screens/HomeScreen';
import DetailsScreen from './screens/DetailsScreen';

const AppStackNavigator = createStackNavigator();

const App = () => {
  return (
    <NavigationContainer>
      <AppStackNavigator.Navigator initialRouteName="Home">
        <AppStackNavigator.Screen name="Home" component={HomeScreen} options={{ title: 'Home' }} />
        <AppStackNavigator.Screen name="Details" component={DetailsScreen} options={{ title: 'Details' }} />
      </AppStackNavigator.Navigator>
    </NavigationContainer>
  );
};

export default App;