/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import * as React from 'react';

import { Close as CloseIcon } from '@mui/icons-material';
import { Button, IconButton, InputAdornment, Tooltip } from '@mui/material';

import { CommandPaletteContext } from './command-palette-context';

export const CommandPaletteSuffix = React.memo(() => {
  const { inputValue, setInputValue, cta } = React.useContext(CommandPaletteContext);

  const onClear = React.useCallback(() => {
    setInputValue('');
  }, [setInputValue]);

  return (
    <InputAdornment position='end' variant='filled'>
      {inputValue.length > 0 && (
        <IconButton onClick={onClear}><CloseIcon /></IconButton>
      )}
      { cta && (
        <Tooltip title={cta.tooltip}>
          <span> {/* Required for the tooltip to show up when the button is disabled */}
            <Button
              size='small'
              variant='contained'
              disabled={cta.disabled === true}
              // eslint-disable-next-line react-memo/require-usememo
              sx={{
                // Push the edges of the button to the edges of the input, and square its left side
                p: (t) => t.spacing(1),
                mr: (t) => `calc(${t.spacing(-3 / 4)})`,
                borderTopLeftRadius: 0,
                borderBottomLeftRadius: 0,
              }}
              onClick={cta.action}
            >
                {cta.label}
            </Button>
          </span>
        </Tooltip>
      )}
    </InputAdornment>
  );
});
CommandPaletteSuffix.displayName = 'CommandPaletteSuffix';
