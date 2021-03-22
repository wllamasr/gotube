import React, {useState} from 'react';
import {Breadcrumb, Layout, Menu} from 'antd';
import {Routes} from './Routes'
import './App.css'

import {
    FileOutlined,
    LaptopOutlined,
    MenuFoldOutlined,
    MenuUnfoldOutlined,
    NotificationOutlined,
    UserOutlined,
} from '@ant-design/icons';

const {Header, Content, Footer, Sider} = Layout;
const {SubMenu} = Menu;

function App() {
    const [collapsed, setCollapsed] = useState<boolean>(true)

    const onCollapse = () => {
        setCollapsed(!collapsed)
    }

    return <Layout style={{minHeight: '100vh'}}>
        <Header className="header">
            <div className="logo"/>
            <Menu theme="dark" mode="horizontal" defaultSelectedKeys={['2']}>
                {React.createElement(collapsed ? MenuUnfoldOutlined : MenuFoldOutlined, {
                    className: 'trigger',
                    onClick: onCollapse,
                })}
                <Menu.Item key="1">nav 1</Menu.Item>
                <Menu.Item key="2">nav 2</Menu.Item>
                <Menu.Item key="3">nav 3</Menu.Item>
            </Menu>
        </Header>
        <Layout>
            <Sider trigger={null} width={200} className="site-layout-background" collapsible collapsed={collapsed}
                   onCollapse={onCollapse}>
                <Menu
                    mode="inline"
                    style={{height: '100%', borderRight: 0}}
                >
                    <SubMenu key="sub1" icon={<UserOutlined/>} title="subnav 1">
                        <Menu.Item key="1">option1</Menu.Item>
                        <Menu.Item key="2">option2</Menu.Item>
                        <Menu.Item key="3">option3</Menu.Item>
                        <Menu.Item key="4">option4</Menu.Item>
                    </SubMenu>
                    <SubMenu key="sub2" icon={<LaptopOutlined/>} title="subnav 2">
                        <Menu.Item key="5">option5</Menu.Item>
                        <Menu.Item key="6">option6</Menu.Item>
                        <Menu.Item key="7">option7</Menu.Item>
                        <Menu.Item key="8">option8</Menu.Item>
                    </SubMenu>
                    <SubMenu key="sub3" icon={<NotificationOutlined/>} title="subnav 3">
                        <Menu.Item key="9">option9</Menu.Item>
                        <Menu.Item key="10">option10</Menu.Item>
                        <Menu.Item key="11">option11</Menu.Item>
                        <Menu.Item key="12">option12</Menu.Item>
                    </SubMenu>
                    <Menu.Item key="9" icon={<FileOutlined/>}>
                        Files
                    </Menu.Item>
                </Menu>
            </Sider>
            <Layout style={{padding: '0 24px 24px'}}>
                <Breadcrumb style={{margin: '16px 0'}}>
                    <Breadcrumb.Item>Home</Breadcrumb.Item>
                    <Breadcrumb.Item>List</Breadcrumb.Item>
                    <Breadcrumb.Item>App</Breadcrumb.Item>
                </Breadcrumb>
                <Content
                    className="site-layout-background"
                    style={{
                        padding: 24,
                        margin: 0,
                        minHeight: 280,
                    }}
                >
                    <Routes/>
                </Content>
                <Footer style={{textAlign: 'center'}}>GoTube - A Slow development process</Footer>
            </Layout>
        </Layout>
    </Layout>
}

export default App;
