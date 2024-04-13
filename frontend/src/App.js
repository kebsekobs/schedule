import "./App.css";
import Header from "./header/Header";
import Router from "./utils/Router";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

function App() {
    const queryClient = new QueryClient()
  return (
      <QueryClientProvider client={queryClient}>
        <div className="App">
          <Header />
          <Router />
        </div>
      </QueryClientProvider>
  );
}

export default App;
