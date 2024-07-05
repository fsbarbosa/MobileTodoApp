import React from 'react';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import HomeScreen from './screens/HomeScreen';
import DetailsScreen from './screens/DetailsScreen';

const RootStack = createStackNavigator();

const App = () => {
  return (
    <NavigationContainer>
      <RootStack.Navigator initialRouteName="HomeScreen">
        <RootStack.Screen name="HomeScreen" component={HomeScreen} options={{ title: 'Home' }} />
        <RootStack.Screen name="DetailsScreen" component={DetailsScreen} options={{ title: 'Details' }} />
      </RootStack.Navigator>
    </NavigationContainer>
  );
};

export default App;