import { PageContainer, ProColumns, ProTable } from '@ant-design/pro-components';
import React, { useRef, useState } from 'react';
import { FormattedMessage } from 'umi';

const columns: ProColumns<API.TaskList>[] = [
    {
        title: <FormattedMessage id='pages.task.id' defaultMessage='ID'></FormattedMessage>,
        valueType: 'textarea',
    },
    {
        title: <FormattedMessage id='pages.task.content' defaultMessage='Content'></FormattedMessage>,
        valueType: 'textarea',        
    },
    {
        title: <FormattedMessage id='pages.task.pc' defaultMessage='Process Percent'></FormattedMessage>,
        valueType: 'textarea',
    },
    {
        title: <FormattedMessage id='pages.task.status' defaultMessage='Status'></FormattedMessage>,
        valueType: 'textarea',
    },
    {
        title: <FormattedMessage id='pages.task.expiryTime' defaultMessage='ExpiryTime'></FormattedMessage>,
        valueType: 'textarea',
    }
];

const TaskList: React.FC = () => {
    return (
        <PageContainer>
            <ProTable<API.TaskList>
                columns={columns}
            >
            </ProTable>
        </PageContainer>
    );
}

export default TaskList