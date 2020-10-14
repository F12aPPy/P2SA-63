import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import ComponanceTable from '../Table';
import Button from '@material-ui/core/Button';
 
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบฝากครรภ์' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName}`}
       subtitle="นพ.AAA BBBB"
     ></Header>
     <Content>
       <ContentHeader title="ตารางข้อมูลฝากครรภ์">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             Add User
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
