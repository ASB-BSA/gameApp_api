import React from 'react';
import { Route, Routes } from 'react-router-dom';
import RouterConfig from './config/RouterConfig';

function App() {
  return (
    <Routes>
      {RouterConfig.map(e => {
        return (
          <>
            <Route path={e.path ? e.path : ""} element={e.component}>
            {
              e.children?.map(el => {
                return (
                  <>
                    <Route path={el.path ? el.path : ""} element={el.component}>
                      {
                        el.children?.map(elm => {
                          return (
                            <>
                              <Route path={elm.path ? elm.path : ""} element={elm.component} />
                            </>
                          )
                        })
                      }
                    </Route>
                  </>
                )
              })
            }
            </Route>
          </>
        )
      })}
    </Routes>
  );
}

export default App;
