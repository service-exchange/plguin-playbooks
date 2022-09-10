// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import {Store as BaseStore} from 'redux';
import {ThunkDispatch} from 'redux-thunk';
import {GlobalState} from '@service-exchange/types/store';

/**
 * Emulated Store type used in service-exchange-webapp/service-exchange-redux
 */
export type Store = BaseStore<GlobalState> & {dispatch: Dispatch}

export type Dispatch = ThunkDispatch<GlobalState, any, any>
